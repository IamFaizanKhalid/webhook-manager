$(document).ready(function() {
    // Retrieve hooks and display them
    $.get('http://localhost:8000/hooks', function(hooks) {
        let hooksContainer = $('#hooks-container');
        for (let i = 0; i < hooks.length; i++) {
            let hook = hooks[i];
            let card = $('<div>').addClass('col-md-4 border rounded');
            let cardBody = $('<div>').addClass('card-body');
            let cardTitle = $('<h5>').addClass('card-title').text(hook.id);
            let cardText = $('<p>').addClass('card-text').html('<i class="fa fa-folder-o"></i> '+hook.command_working_directory);
            let editButton = $('<a>').addClass('text-info mr-2').attr('href', '/edit_hook?id=' + hook.id).html('<i class="fa fa-pencil"></i>');
            let deleteButton = $('<a>').addClass('text-danger').attr('href', '#').html('<i class="fa fa-trash"></i>');

            // Add click event to delete button
            deleteButton.click(function(event) {
                event.preventDefault();
                $.ajax({
                    url: 'http://localhost:8000/hooks/' + hook.id,
                    type: 'DELETE',
                    success: function() {
                        // Remove card from display
                        card.remove();
                    }
                });
            });

            cardBody.append(cardTitle, cardText, editButton, deleteButton);
            card.append(cardBody);
            hooksContainer.append(card);
        }
    });
});
