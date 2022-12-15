$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(event) {
    event.preventDefault();
    console.log('criarUsuario')

    if ( $('#senha').val() != $('#confirmaSenha').val() ) {
        Swal.fire("Ops...", "As senhas não coincidem!", "error")
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
        Swal.fire("Sucesso!", "Usuário cadastrado!", "success")
            .then(function() {
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $('#email').val(),
                        senha: $('#senha').val()
                    }

                } ).done(function() {
                    window.location = "/home";

                } ).fail(function() {
                    Swal.fire("Ops...", "Erro ao autenticar o usuário!", "error");
                } );
            });
    } ).fail(function() {
        Swal.fire("Ops...", "Erro ao cadastrar usuário!", "error")
    } );
}
