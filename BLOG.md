# My journey with HTMX

*written by Luc B. on the 10th Jan 2024*

HTMX has been quite hyped for a few months.
I have discovered it through Youtubers or articles and it caught my attention because it presented a "new" developing paradigm.
Here, I will try to develop a simple framework built around HTMX that would allow me to build rapidly some easy websites.

## HTMX

HTMX is a lightweight tool that upgrades your HTML by providing attributes that grant AJAX capabilities.
As said on [their website](https://htmx.org/),

> htmx gives you access to AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML, using attributes, so you can build modern user interfaces with the simplicity and power of hypertext

Pretty cool, huh ?! The idea of building websites without using a big JS framework and focusing back on the HTML sounds really interesting and, well, kinda new for someone who got into development after React was created and widespread.

Here is a list of what interested me :

- I am tired of JS frameworks that need their own dev server and that can be really slow
- I also tend to dislike build processes, as they can be long
- Some frameworks or meta-frameworks can be complex and unintuitive
- We can have the frontend built alongside the backend, which simplifies processes since there is no compatibility issues
- Having simply HTML + a lightweight script can boost significantly the performance of web apps

## Side note on JS frameworks

Even though I said I was tired of JS frameworks, it was because of the ecosystem around them. I actually love frameworks like Vue or Svelte because the fact that you can split youe application in simple testable components is amazing !

There are also a lot of tools built (sometimes built-in) for these frameworks that makes your life really easier, like stores, component libraries, etc.

I think frameworks such as React, Vue and others are in a really good place right now and have become the go-to way of writing frontends. And I want to take inspiration from all the good things in them, especially how components are built and combined.

## Framework requirements

To recap what I said, here is a list of the things I want :

- A very simple framework being easy to understand
- No hidden or unintuitive features
- Separation into components
- No template files, templates in the component files
- Ability to write styles and/or scripts alongside the HTML
- No Node.JS
- Static file serving
- Automatic routing

The framework isn't meant to be complete, it is meant to be "good-enough" to build simple websites without too much hurdle and to be fast in its execution, its writing and its development cycles.

## The backend

Interestingly, I gave a lot of thought about which backend to choose.

I initially thought of Express and other JS/TS backend servers after I watched Ethan Niser's video on [the BETH stack](https://www.youtube.com/watch?v=cpzowDDJj24). However, I don't think using JS backends in my context makes sense, because of their performance and complex ecosystem.

I also saw a lot of people using HTMX with Django, which makes a lot of sense. I think they work really well together, but I am not a fan of the templating (and I didn't really want to go with Python).

I finally settled onto Go because it provides good performance, isn't very complicated and has built-in tools for HTTP servers and HTML templating.

## The structure

Based on my requirements, I 
