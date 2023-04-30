
$(document).ready(function() {
    const submitButton = $('#hook-form');

    submitButton.click(function(event) {
        event.preventDefault();

        // Get form values
        const executeCommand = $('#executeCommand').val();
        const commandWorkingDirectory = $('#commandWorkingDirectory').val();
        const responseMessage = $('#responseMessage').val();
        const responseHeaders = $('#responseHeaders').val();
        const captureCommandOutput = $('#captureCommandOutput').prop('checked');
        const captureCommandOutputOnError = $('#captureCommandOutputOnError').prop('checked');
        const passEnvironmentToCommand = $('#passEnvironmentToCommand').val();
        const passArgumentsToCommand = $('#passArgumentsToCommand').val();
        const passFileToCommand = $('#passFileToCommand').val();
        const JSONStringParameters = $('#JSONStringParameters').val();
        const triggerRule = $('#triggerRule').val();
        const triggerRuleMismatchHttpResponseCode = parseInt($('#triggerRuleMismatchHttpResponseCode').val());
        const triggerSignatureSoftFailures = $('#triggerSignatureSoftFailures').prop('checked');
        const incomingPayloadContentType = $('#incomingPayloadContentType').val();
        const successHttpResponseCode = parseInt($('#successHttpResponseCode').val());
        const httpMethods = $('input[name="httpMethods"]:checked').map(function() {
            return $(this).val();
        }).get();

        // Construct payload object
        const payload = {
            execute_command: executeCommand,
            command_working_directory: commandWorkingDirectory,
            response_message: responseMessage,
            response_headers: responseHeaders,
            include_command_output_in_response: captureCommandOutput,
            include_command_output_in_response_on_error: captureCommandOutputOnError,
            pass_environment_to_command: passEnvironmentToCommand,
            pass_arguments_to_command: passArgumentsToCommand,
            pass_file_to_command: passFileToCommand,
            parse_parameters_as_json: JSONStringParameters,
            trigger_rule: triggerRule,
            trigger_rule_mismatch_http_response_code: triggerRuleMismatchHttpResponseCode,
            trigger_signature_soft_failures: triggerSignatureSoftFailures,
            incoming_payload_content_type: incomingPayloadContentType,
            success_http_response_code: successHttpResponseCode,
            http_methods: httpMethods
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