$('#parar-de-seguir').on('click', pararDeSeguir)
$('#seguir').on('click', seguir)

function pararDeSeguir( event ) {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/parar-de-seguir`,
        method: "POST"
        
    } ).done( function() {
        document.location = `/usuarios/${usuarioId}`;

    } ).fail( function() {
        Swal.fire( "Ops...", "Erro ao parar de seguir o usuário", "error" );
        $(this).prop( 'disabled', false );
    } );
}

function seguir( event ) {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/usuarios/${usuarioId}/seguir`,
        method: "POST"
        
    } ).done( function() {
        document.location = `/usuarios/${usuarioId}`;

    } ).fail( function() {
        Swal.fire( "Ops...", "Erro ao seguir o usuário", "error" );
        $(this).prop( 'disabled', false );
    } );
}