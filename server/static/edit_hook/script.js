$(document).ready(function () {
    // Check if API key exists in local storage
    var apiKey = localStorage.getItem('api_key');
    if (!apiKey) {
        window.location.href = '/login';
        return
    }

    // Get the ID from query params
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');
    if (id == null || id === '') {
        window.location.href = '/';
        return
    }

    // Make GET request to the API endpoint
    $.ajax({
        url: '/hooks/'+id,
        type: 'GET',
        headers: {"X-API-KEY": apiKey},
        success: function(hook) {
            // Fill the form with the data
            $('#id').val(hook.id);
            $('#cmd').val(hook.execute_command);
            $('#pwd').val(hook.command_working_directory);
            $('#branch').val(hook.trigger_rule.and[1].match.value.replace('refs/heads/', ''));
            $('#secret').val(hook.trigger_rule.and[0].match.secret);
            $('#hook-url').text(location.protocol + '//' + location.hostname + ':9000/hooks/'+id);
        },
        error: function (error) {
            console.error(error.responseJSON.error);
        }
    });



    $("#togglePassword").click(function() {
        var password = $("#secret");
        var icon = $(this).find("i");
        if (password.attr("type") === "password") {
            password.attr("type", "text");
            icon.removeClass("fa-eye").addClass("fa-eye-slash");
        } else {
            password.attr("type", "password");
            icon.removeClass("fa-eye-slash").addClass("fa-eye");
        }
    });

    $('#edit-hook').click(function (event) {
        event.preventDefault();

        // Construct payload object
        const payload = {
            id: $('#id').val(),
            execute_command: $('#cmd').val(),
            command_working_directory: $('#pwd').val(),
            trigger_rule: {
                and: [
                    {
                        match: {
                            type: "payload-hash-sha256",
                            secret: $('#secret').val(),
                            parameter: {
                                source: "header",
                                name: "X-Hub-Signature-256"
                            }
                        }
                    },
                    {
                        match: {
                            type: "value",
                            value: "refs/heads/" + $('#branch').val(),
                            parameter: {
                                source: "payload",
                                name: "ref"
                            }
                        }
                    }
                ]
            },
            pass_arguments_to_command: [
                {
                    source: "payload",
                    name: "head_commit.id"
                },
                {
                    source: "payload",
                    name: "head_commit.message"
                },
                {
                    source: "payload",
                    name: "head_commit.author.name"
                },
                {
                    source: "payload",
                    name: "head_commit.author.email"
                }
            ]
        };

        // Send POST request to API endpoint
        $.ajax({
            url: '/hooks/'+id,
            type: 'PUT',
            dataType: 'json',
            contentType: 'application/json',
            headers: {"X-API-KEY": apiKey},
            data: JSON.stringify(payload),
            success: function (data) {
                window.location.href = '/';
            },
            error: function (error) {
                $('#error-message').text(error.responseJSON.error).css('display','inline-block');
            }
        });
    });

    $('#logout').click(function (event) {
        event.preventDefault();

        localStorage.removeItem('api_key');
        window.location.href = '/login';
    });
});
