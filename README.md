<a name="readme-top"></a>
# GO-Reepay
Golang integration of [Reepay API](https://reference.reepay.com/api/#introduction "Official API reference from reepay.")

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project
This is a project made to integrate with Reepay to handle payments for systems such as a webshop.

Example API is located in the /example folder. Made with GIN

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* make sure go is installed
[https://go.dev/doc/install](install golang)

### Installation


1. Clone the repo
   ```sh
   git clone https://github.com/PhilipGeil/go-reepay
   ```
2. Install GO packages
   ```sh
   go get .
   ```
3. Enter your key and url's in `.env`
   ```dotenv
    REEPAY_PRIV_KEY='ENTER YOUR API'
    ACCEPT_URL='enter your accept url'
    CANCEL_URL='enter your cancel url'
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- USAGE EXAMPLES -->
## Usage

Run the program located in /example and the API will be available on port 8080 as default.

[//]: # (_For more examples, please refer to the [Documentation]&#40;https://example.com&#41;_)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [ ] Support charge
  - [x] Create charge
  - [x] Settle charge
  - [x] Refund charge
  - [x] Cancel charge
  - [x] Get charge
  - [ ] Get list of charges
- [ ] Support subscription
- [x] Add license
- [ ] Set up tests
- [ ] Documentation

See the [open issues](https://github.com/PhilipGeil/go-reepay/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>