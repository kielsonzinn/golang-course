$('#nova-publicacao').on('submit', criarPublicacao);

$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descutir-publicacao', descurtirPublicacao);

$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);

function criarPublicacao( event ) {
    event.preventDefault();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function() {
        window.location = "/home"
    }).fail(function() {
        Swal.fire( "Ops...", "Erro ao criar publicação!", "error" );
    });
}

function curtirPublicacao( event ) {
    event.preventDefault();

    const elementClicked = $(event.target);
    const publicacaoId = elementClicked.closest( 'div' ).data( 'publicacao-id' );

    elementClicked.prop('disabled', true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"
    }).done(function() {
        const contadorDeCurtidas = elementClicked.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());
        contadorDeCurtidas.text(quantidadeDeCurtidas + 1);

        elementClicked.addClass('descutir-publicacao');
        elementClicked.addClass('text-danger');
        elementClicked.removeClass('curtir-publicacao');

    }).fail(function() {
        Swal.fire( "Ops...", "Erro ao curtir publicação!", "error" );

    }).always(function() {
        elementClicked.prop('disabled', false);
    });
}

function descurtirPublicacao( event ) {
    event.preventDefault();

    const elementClicked = $(event.target);
    const publicacaoId = elementClicked.closest( 'div' ).data( 'publicacao-id' );

    elementClicked.prop('disabled', true);
    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST"
    }).done(function() {
        const contadorDeCurtidas = elementClicked.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());
        contadorDeCurtidas.text(quantidadeDeCurtidas - 1);

        elementClicked.removeClass('descutir-publicacao');
        elementClicked.removeClass('text-danger');
        elementClicked.addClass('curtir-publicacao');

    }).fail(function() {
        Swal.fire( "Ops...", "Erro ao curtir publicação!", "error" );

    }).always(function() {
        elementClicked.prop('disabled', false);
    });
}

function atualizarPublicacao(event) {
    $(this).prop('disabled', true);

    const publicacaoId = $(this).data('publicacao-id');
    
    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function() {
        Swal.fire( "Sucesso!", "Publicação atualizada com sucesso!", "success" )
            .then(function() {
                window.location = "/home";
            });

    }).fail(function() {
        Swal.fire( "Ops...", "Erro ao atualizar a publicação", "error" );

    }).always(function() {
        $('#atualizar-publicacao').prop('disabled', false);
    });

}

function deletarPublicacao( event ) {
    event.preventDefault();

    Swal.fire( {
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa publicação? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButton: "Cancelar",
        icon: "warning"
    } ).then( function( confirm ) {
        if ( !confirm.value ) {
            return;
        }

        const elementClicked = $(event.target);
        const publicacao = elementClicked.closest( 'div' );
        const publicacaoId = publicacao.data( 'publicacao-id' );

        elementClicked.prop('disabled', true);

        $.ajax({
            url: `/publicacoes/${publicacaoId}`,
            method: "DELETE"
        }).done(function() {
            publicacao.fadeOut("slot", function() {
                $(this).remove();
            })

        } ).fail( function() {
            Swal.fire( "Ops...", "Erro ao deletar publicação!", "error" );

        } ).always( function() {
            elementClicked.prop('disabled', false);
        } );
    } );
}