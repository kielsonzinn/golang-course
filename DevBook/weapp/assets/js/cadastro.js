$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(event) {
    event.preventDefault();
    console.log('criarUsuario')

    if ( $('#senha').val() != $('#confirmaSenha').val() ) {
        alert('As senhas não coincidem!')
        return;
    }

    $.ajax( {
        url: '/usuarios',
        method: 'POST',
        data:  {
            nome: $('#nome').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            senha: $('#senha').val()
        }
    } ).done(function() {
        alert("Usuário cadastrado com sucesso!");
    }).fail(function() {
        alert("Erro ao cadastrar usuário!");
    });
}
