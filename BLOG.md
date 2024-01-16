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

- I am tired of JS frontend frameworks that need their own dev server and that can be really slow
- I also tend to dislike build processes, as they can be long
- Some frameworks or meta-frameworks can be complex and unintuitive
- We can have the frontend built alongside the backend, which simplifies processes since there is no compatibility issues
- Having simply HTML + a lightweight script can boost significantly the performance of web apps

## Side note on JS frameworks

Even though I said I was tired of JS frameworks, it was because of the ecosystem around them. I actually love frameworks like Vue or Svelte because the fact that you can split your application in simple testable components is amazing !

There are also a lot of tools built (sometimes built-in) for these frameworks that makes your life really easier, like stores, component libraries, etc.

I think frameworks such as React, Vue and others are in a really good place right now and have become the go-to way of writing frontends. And I want to take inspiration from all the good things in them, especially how components are built and combined.

## Framework requirements

To recap what I said, here is a list of the things I want :

- A very simple framework being easy to understand
- No hidden or unintuitive features
- Separation into components
- No separate template files, templates in the component files
- Easy-to-write and understand templates (via string interpolation)
- Ability to write styles and/or scripts alongside the HTML
- Static file serving
- Automatic routing

The framework isn't meant to be complete, it is meant to be "good-enough" to build simple websites without too much hurdle and to be fast in its execution, its writing and its development cycles.

## The backend

I gave a lot of thought about which backend to choose, because this is where most of the work will be.

Here are the options I considered :

- Express JS : very popular, very inuitive, but performance might be an issue
- Elysia: the one used in Ethan Niser's video on [the BETH stack](https://www.youtube.com/watch?v=cpzowDDJj24), with JSX support, but not top-of-the-pack performance
- Django: also very popular, especially alongside HTMX, but I don't really like the templating and performance
- Go : very good performance, interesting language with many built-in tools, but unintuitive templating IMO.
- C# / .NET Core : good performance, great string interpolation, but might be a bit complex for simple applications.

Based on this quick comparison and even though I wanted to distance myself from JS/TS for this project, I decided to try Elysia for its easy setup and JSX support.

## The structure

Based on my requirements, I have a few folders and files in mind to create. Here is a non-exhaustive list :

- `routes` : Folder where we will place html files that will be automatically served. These html files are the 'base' of our pages onto which we will incorporate our components
- `components` : Folder where will will have our components
- `public` : Static files to be served
- `src` : Folder where I will put the custom code for the framework (may be placed in a library in the future)
