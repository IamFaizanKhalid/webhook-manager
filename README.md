# webhook-manager

Configure automatic deployments easily using GitHub Webhooks.

<img align="right" width="250" alt="Webhook Manager" src="https://github.com/IamFaizanKhalid/webhook-manager/raw/master/server/static/logo.png">


## Pre-requisites
- [webhook](https://github.com/adnanh/webhook)


## Setup

- Create `.env` file like `.env.example`.
- Set `API_KEY` variable in `.env` to use as a password.
- Run the following command to set up this application
    ```
    ./scripts/setup.sh
    ```

## Usage

To use _webhook-manager_ visit port `:8000` of your server, e.g. http://www.example.com:8000.

To configure auto deployment for your application follow the following steps

- Write a script which perform all the operations required for the deployment of your application.
- Add a new hook in _webhook-manager_.
- Note down your payload **URL** and **secret**.
- Go to settings of your repository in the Webhooks section. \
https://github.com/{user}/{repo}/settings/hooks
- Add a new webhook and configure it.
- All set..!
