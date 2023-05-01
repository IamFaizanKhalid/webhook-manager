$(document).ready(function () {
    // Check if API key exists in local storage
    const apiKey = localStorage.getItem('api_key');

    if (apiKey && window.location.pathname === '/login/') {
        window.location.href = '/';
        return
    } else if (!apiKey && window.location.pathname !== '/login/') {
        window.location.href = '/login';
        return
    }

    // Handle logout
    $('#logout').click(function (event) {
        event.preventDefault();

        localStorage.removeItem('api_key');
        window.location.href = '/login';
    });

    // Get the ID from query params
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');

    function submit_form(method, url) {
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
            url: url,
            type: method,
            dataType: 'json',
            contentType: 'application/json',
            headers: {"X-API-KEY": apiKey},
            data: JSON.stringify(payload),
            success: function (data) {
                window.location.href = '/';
            },
            error: function (error) {
                $('#error-message').text(error.responseJSON.error).css('display', 'inline-block');
            }
        });
    }

    // page specific
    switch (window.location.pathname) {
        // homepage
        case '/':
            // Retrieve hooks and display them
            $.ajax({
                url: '/hooks',
                type: 'GET',
                headers: {"X-API-KEY": apiKey},
                success: function (hooks) {
                    let hooksContainer = $('#hooks-container');
                    for (let i = 0; i < hooks.length; i++) {
                        let hook = hooks[i];
                        let project = hook.id;
                        let payloadUrl = location.protocol + '//' + location.hostname + ":9000/hooks/" + project;
                        let branch = hook.trigger_rule.and[1].match.value.replace('refs/heads/', '');

                        let div = $('<div>').addClass('card col-md-4 border rounded mb-3 shadow-sm');
                        let divBody = $('<div>').addClass('card-body');
                        let divTitle = $('<h5>').addClass('card-title').text(project);
                        let folderIcon = $('<i>').addClass('fa fa-globe mr-2');
                        let branchIcon = $('<i>').addClass('fa fa-code-fork mr-2');
                        let directoryText = $('<p>').addClass('card-text').append(folderIcon).append(payloadUrl);
                        let branchText = $('<p>').addClass('card-subtitle mb-2 text-muted').append(branchIcon).append(branch);
                        let editLink = $('<a>').addClass('card-link text-info mr-2').attr('href', '/edit_hook?id=' + hook.id).html('<i class="fa fa-pencil"></i>');
                        let deleteLink = $('<a>').addClass('card-link text-danger').attr('href', '#').html('<i class="fa fa-trash"></i>');

                        // Add click event to delete button
                        deleteLink.click(function (event) {
                            event.preventDefault();
                            let confirmationModal = $('#confirmation-modal');
                            let confirmationButton = $('#confirmation-button');

                            // Set the hook id in the confirmation button data attribute
                            confirmationButton.attr('data-hook-id', hook.id);

                            // Show the confirmation modal
                            confirmationModal.modal('show');
                        });

                        // Add click event to confirmation button
                        $('#confirmation-button').click(function (event) {
                            event.preventDefault();
                            let hookId = $(this).attr('data-hook-id');
                            $.ajax({
                                url: '/hooks/' + hookId,
                                type: 'DELETE',
                                headers: {"X-API-KEY": apiKey},
                                success: function () {
                                    // Remove div from display
                                    div.remove();
                                },
                                error: function (error) {
                                    console.error(error.responseJSON.error);
                                }
                            });

                            // Hide the confirmation modal
                            $('#confirmation-modal').modal('hide');
                        });


                        divBody.append(divTitle, branchText, directoryText, editLink, deleteLink);
                        div.append(divBody);
                        hooksContainer.append(div);
                    }
                },
                error: function (error) {
                    console.error(error.responseJSON.error);
                }
            });

            break;

        // add hook page
        case '/add_hook/':
            $('#hook-url').text(location.protocol + '//' + location.hostname + ":9000/hooks/my-app");

            $('#add-hook').click(function (event) {
                event.preventDefault();

                submit_form('POST', '/hooks');
            });

            break;

        // edit hook page
        case '/edit_hook/':
            if (id == null || id === '') {
                window.location.href = '/';
                return
            }

            // Make GET request to the API endpoint
            $.ajax({
                url: '/hooks/' + id,
                type: 'GET',
                headers: {"X-API-KEY": apiKey},
                success: function (hook) {
                    // Fill the form with the data
                    $('#id').val(hook.id);
                    $('#cmd').val(hook.execute_command);
                    $('#pwd').val(hook.command_working_directory);
                    $('#branch').val(hook.trigger_rule.and[1].match.value.replace('refs/heads/', ''));
                    $('#secret').val(hook.trigger_rule.and[0].match.secret);
                    $('#hook-url').text(location.protocol + '//' + location.hostname + ':9000/hooks/' + id);
                },
                error: function (error) {
                    console.error(error.responseJSON.error);
                }
            });

            $('#edit-hook').click(function (event) {
                event.preventDefault();

                submit_form('PUT', '/hooks/' + id);
            });

            break;

        // edit hook page
        case '/login/':

            $('#login').click(function (event) {
                event.preventDefault();

                const apiKey = $('#secret').val();

                // Send password to /hello endpoint with API key in header
                $.ajax({
                    url: '/hello',
                    type: 'GET',
                    headers: {'X-API-KEY': apiKey},
                    success: function (data) {
                        // Redirect to home page on success
                        localStorage.setItem('api_key', apiKey);
                        window.location.href = '/';
                    },
                    error: function () {
                        // Display error message on failure
                        $('#error-message').text('Invalid password').css('display', 'inline-block');
                    }
                });
            });

            break;

        // unknown page
        default:
            window.location.href = '/';
            return
    }


    $("#togglePassword").click(function () {
        let password = $("#secret");
        let icon = $(this).find("i");
        if (password.attr("type") === "password") {
            password.attr("type", "text");
            icon.removeClass("fa-eye").addClass("fa-eye-slash");
        } else {
            password.attr("type", "password");
            icon.removeClass("fa-eye-slash").addClass("fa-eye");
        }
    });


});
