$(document).ready(function () {
    // Check if API key exists in local storage
    var apiKey = localStorage.getItem('api_key');
    if (apiKey) {
        window.location.href = '/';
        return
    }

    // Handle form submission
    $('form').submit(function(event) {
        event.preventDefault();
        var apiKey = $('input[name=password]').val();

        // Send password to /validate endpoint with API key in header
        $.ajax({
            url: '/hello',
            type: 'GET',
            headers: { 'X-API-KEY': apiKey },
            success: function(data) {
                // Redirect to home page on success
                localStorage.setItem('api_key',apiKey);
                window.location.href = '/';
            },
            error: function() {
                // Display error message on failure
                $('#error-message').text('Invalid password').css('display','inline-block');
            }
        });
    });


    $("#togglePassword").click(function() {
        var password = $("#password");
        var icon = $(this).find("i");
        if (password.attr("type") === "password") {
            password.attr("type", "text");
            icon.removeClass("fa-eye").addClass("fa-eye-slash");
        } else {
            password.attr("type", "password");
            icon.removeClass("fa-eye-slash").addClass("fa-eye");
        }
    });
});
