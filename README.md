<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="Showcase">
    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/8/84/Spotify_icon.svg/512px-Spotify_icon.svg.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Spotify Account Generator</h3>

  <p align="center">
    Create thousands of accounts in minutes!
    <br />
        <a href="https://www.youtube.com/watch?v=Jm9v0CYITUc"><strong>Showcase »</strong></a>
    <br />
  </p>
</div>


<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://spotium.store)

With all the spotify account generators out there, they all suck! This solves all those problems.

# Here's why:
* Have not been updated in a while
* Lack scalability and error handling
* Uses old methods of creating accounts

# Why did I release this?
* I want a portable library for my Spotify Auto Upgrader bot to clear up some work space
* I found no use in keeping it private
* I'd like to grow my profile a bit

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Golang][Golang]][Golang-Url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

This works for all operating systems as long as you have the requirements installed.

### Prerequisites

Golang is a must (Version 1.19) 

### Installation

_If you are on windows and do not care about running the program from source, you can find the exe in /Bin. Just drag it out and run it, and you are good to go._

1. Fill out your Config will all the required information. The captcha client supported is https://capsolver.com/. 
2. Clone the repo, or download it from github.
   ```sh
   git clone https://github.com/Aran404/Spotify-Account-Creator
   ```
3. Run the main file.
   ```sh
   go run .
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

To **Avoid** rate limit, you must use HTTP proxies. The ones I used in the video were [SnowProxies](https://snowproxies.digital/).
The issue with some proxies is that they will cause the "Challenge" captcha to run meaning you'll need to solve a RecaptchaV2 challenge.
Depending on your proxies you'll not need to solve the second captcha. You can figure out which way is cheaper for you.
In my case they were insanely cheap due to IPV6. I highly recommend [SnowProxies](https://snowproxies.digital/).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Discord - **@optiboosts**\n
Cheap Spotify Premium - [Website](https://spotium.store/) | [Discord](https://discord.gg/Spotium)


<p align="right">(<a href="#readme-top">back to top</a>)</p>




<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[forks-shield]: https://img.shields.io/github/forks/Aran404/Spotify-Account-Generator.svg?style=for-the-badge
[forks-url]: https://github.com/Aran404/Spotify-Account-Generator/network/members
[stars-shield]: https://img.shields.io/github/stars/Aran404/Spotify-Account-Generator.svg?style=for-the-badge
[stars-url]: https://github.com/Aran404/Spotify-Account-Generator/stargazers
[product-screenshot]: https://media.discordapp.net/attachments/1154772032551141507/1161826422969602149/image.png?ex=6539b62c&is=6527412c&hm=a879e01ac9891ff56da3fd850f19d4610bb42b8030e459f6d4e6c2de18a8509d&=
[Golang]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[Golang-Url]: https://go.dev/
