## Email-Dispatcher (Go): Bulk Mail with Just One Command

A bulk email dispatcher written in Go using SMTP, CSV, and templating with concurrency.

## Features

- Load recipients from CSV
- Personalized email templates
- Email body from `.txt` file
- CLI-based subject input
- SMTP-based sending (Gmail supported)
- Worker pool concurrency
- Env-based SMTP config

## Requirements

- Go 1.25+
- Gmail App Password (if using Gmail)
- `.env` configuration file

## Setup

Clone repository and create `.env` file and add `.csv` file in the same directory/folder

- Example .env file
  ```.env
    SMTP_HOST=smtp.gmail.com
    SMTP_PORT=587
    SMTP_EMAIL=your@gmail.com
    SMTP_PASS=your_app_password
    SMTP_FROM_NAME="Your Name"
  ```

- Example .csv file
  ```.csv
    Name,Email
    user1,user1@gmail.com
    user2,user2@gmail.com
  ```

## Email Template

Create a `email.templ` file:

  ```txt
    To: {{.Email}}
    Subject: {{.Subject}}

    Hi {{.Name}},

    {{.Body}}

    Thanks,
    {{.FromName}}
  ```

## Email Body

Create a `body.txt` file and add your email content in it

## Usage

- Cloned the repository for first time you have to run a command to install dependencies:

    ```bash
        go mod tidy
    ```

- Run Command:

    ```bash
        go run . --subject "Your email subject" --body-file "Path to your body.txt file"
    ```

## Gmail Limits

- Free Gmail Accounts: ~500 emails/day
- Workspace Accounts: ~2000 emails/day

## Project Structure

```bash
├── main.go
├── producer.go
├── consumer.go
├── emails.csv
├── email.tmpl
├── body.txt
├── .env
└── README.md
```

## Notes

- Gmail requires App Passwords when 2FA is enabled.
- Concurrency can be configured in `main.go` using `workerCount`.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.