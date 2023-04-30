
$(document).ready(function() {
    const submitButton = $('#hook-form');

    submitButton.click(function(event) {
        event.preventDefault();

        // Construct payload object
        const payload = {
            id: $('#id').val(),
            execute_command: $('#execute-command').val(),
            command_working_directory: $('#command-working-directory').val(),
            response_message: $('#response-message').val(),
            //response_headers: $('#response-headers').val(),
            // capture_command_output: $('#capture-command-output').prop('checked'),
            // capture_command_output_on_error: $('#capture-command-output-on-error').prop('checked'),
            // pass_environment_to_command: $('#pass-environment-to-command').val(),
            // pass_arguments_to_command: $('#pass-arguments-to-command').val(),
            // pass_file_to_command: $('#pass-file-to-command').val(),
            // json_string_parameters: $('#json-string-parameters').val(),
            // trigger_rule: $('#trigger-rule').val(),
            trigger_rule_mismatch_http_response_code: parseInt($('#trigger-rule-mismatch-http-response-code').val()),
            trigger_signature_soft_failures: $('#trigger-signature-soft-failures').prop('checked'),
            incoming_payload_content_type: $('#incoming-payload-content-type').val(),
            success_http_response_code: parseInt($('#success-http-response-code').val()),
            // http_methods: $('input[name="http-methods"]:checked').map(function() {
            //     return $(this).val();
            // }).get()
        };

        // Send POST request to API endpoint
        $.ajax({
            url: 'http://localhost:8000/hooks',
            type: 'POST',
            dataType: 'json',
            contentType: 'application/json',
            data: JSON.stringify(payload),
            success: function(data) {
                console.log(data);
                window.location.href = '/';
            },
            error: function(error) {
                console.error(error);
            }
        });
    });
});