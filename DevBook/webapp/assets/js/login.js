$('#login').on('submit', efetuarLogin);

function efetuarLogin(event) {
    event.preventDefault();
    console.log('efetuarLogin')

    $.ajax( {
        url: '/login',
        method: 'POST',
        data:  {
            email: $('#email').val(),
            senha: $('#senha').val()
        }
    } ).done(function() {
        window.location = "/home";
    }).fail(function() {
        alert("Usuário ou senha incorretos!");
    });
}
