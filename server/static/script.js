$(document).ready(function() {
    // Retrieve hooks and display them
    $.get('http://localhost:8000/hooks', function(hooks) {
        let hooksContainer = $('#hooks-container');
        for (let i = 0; i < hooks.length; i++) {
            let hook = hooks[i];
            let card = $('<div>').addClass('col-md-4 border rounded');
            let cardBody = $('<div>').addClass('card-body');
            let cardTitle = $('<h5>').addClass('card-title').text(hook.id);
            let cardText = $('<p>').addClass('card-text').text(hook.command_working_directory);
            let editButton = $('<a>').addClass('btn btn-primary mr-2').attr('href', '/edit_hook?id=' + hook.id).text('Edit');
            let deleteButton = $('<a>').addClass('btn btn-danger').attr('href', '#').text('Delete');

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
