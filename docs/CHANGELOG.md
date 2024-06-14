# Changelog

## [1.0.0] - 2024-06-14

### Added

- Completion of the project.
- Added and restructured server.
- Added and restructured controller and service layers.
- Added mocks for service layers and server.
- Added test cases.
- Added postman docs.

### Removed

- Removed unneeded directories.
- Removed test dir and created tests within the layer.
- Removed models and initialized the models within layers for more control over testing.

### Further Improvements

- Could have added a `coverage report` and `mem profile` but I am exhausted at this point and the time is about to be
  over.
- Could have added the logger but was short on time.
- The service layer func `AnalysisCalculation` can be restructured such that it follows the options pattern, but I have
  to think about the space complexity of it.
- Could have improved the api test server and the controller initialization. Probably we can reduce some functions and
  make it more simple.
- Could have added test cases for mocks as well as another service that would take variable `Expiration_Times` but that
  was out of scope of documentation.
- `Underlying_Price` was not part of the request object, one idea was to simply init a `Token` in `OptionsContracts` but
  I figured why not just create a service, but it does seem like an overkill.
- It would have been much better to have a `migrations` dir to add the test token, but I was already short on time.
- Code comments could have been better and more explanatory.
- Could have added validation test cases for the correct formulas and value calculations.

### Reference
Reference for formulas was taken from [here](https://analystprep.com/blog/options-calculations-payoff-profit/). Special thanks to the author for a wonderful article.