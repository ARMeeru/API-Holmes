# API Holmes
A simple go program comparing sample response and actual response in bytes. Simply run the binary followed by config JSON path and that's it!

    ./compareAPIresponse test.json
Provide your JSON path accordingly.
# Sample JSON Format

    {
        "api_info": {
            "url": "",
            "method": "",
            "headers": {},
            "request_body": {},
            "sample_response": {}
        }
    }
# TO-DO
 - [ ] Continuation
 - [ ] URL Params Separation Acceptance
 - [ ] Removing Hardcoded Status Code
 - [ ] Better Error Handlings
# Conclusion
This is an attempt at my ongoing golang learning phase. I am eternally grateful to @Anondo

Feel free to criticize, and suggest in the [Discussions](https://github.com/ARMeeru/compareAPIresponse/discussions) tab.

Licensed under [MIT License](https://github.com/ARMeeru/compareAPIresponse/blob/main/LICENSE)
