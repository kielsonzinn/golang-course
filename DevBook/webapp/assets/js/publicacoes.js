$('#nova-publicacao').on('submit', criarPublicacao);

$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descutir-publicacao', descurtirPublicacao);

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
        alert("Erro ao criar publicação");
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
        alert("Erro ao curtir publicação!");

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
        alert("Erro ao curtir publicação!");

    }).always(function() {
        elementClicked.prop('disabled', false);
    });
}