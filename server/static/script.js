$(document).ready(function () {
    // Check if API key exists in local storage
    var apiKey = localStorage.getItem('api_key');
    if (!apiKey) {
        window.location.href = '/login';
        return
    }

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
                let directory = hook.command_working_directory;
                let branch = hook.trigger_rule.and[1].match.value.replace('refs/heads/', '');

                let div = $('<div>').addClass('card col-md-4 border rounded mb-3 shadow-sm');
                let divBody = $('<div>').addClass('card-body');
                let divTitle = $('<h5>').addClass('card-title').text(project);
                let folderIcon = $('<i>').addClass('fa fa-folder mr-2');
                let branchIcon = $('<i>').addClass('fa fa-code-fork mr-2');
                let directoryText = $('<p>').addClass('card-text').append(folderIcon).append(directory);
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

    $('#logout').click(function (event) {
        event.preventDefault();

        localStorage.removeItem('api_key');
        window.location.href = '/login';
    });
});
