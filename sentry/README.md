References:
- Product docs: https://docs.sentry.io/?original_referrer=https%3A%2F%2Fwww.google.com%2F
- Developer docs: https://develop.sentry.dev/
- Relay rust crates: https://getsentry.github.io/relay/relay/
- [Misc] Symbolicator: https://getsentry.github.io/symbolicator/
- Demo Forked from: https://github.com/Ramhm/sentry-setup

![sentry-high-lvl-overview.svg](sentry-high-lvl-overview.svg)


### Instructions

**NOTE**: This setup is purely for exploratory purposes.
- For small scale production usecase: https://github.com/getsentry/self-hosted
- For larger scales: https://github.com/sentry-kubernetes/charts

Step A: Setup sentry
- Refer: https://github.com/getsentry/self-hosted#using-linux

Step B: Test run  & Tinker
- Create a project
- Copy the DSN & integrate with the SDK
- Run the script & watch the event getting recorded into sentry
- Explore other features

Step C: Profit!!!