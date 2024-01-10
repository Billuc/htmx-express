# My journey with HTMX

*written by Luc B. on the 10th Jan 2024*

HTMX has been quite hyped for a few months.
I have discovered it through Youtubers or articles and it caught my attention.
I was interested by this "new" developing paradigm... that looks like old-fashioned HTML.
But what is HTMX ?

## HTMX

HTMX is a lightweight tool that upgrades your HTML by providing attributes that grant AJAX capabilities.
As said on [their website](https://htmx.org/),

> htmx gives you access to AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML, using attributes, so you can build modern user interfaces with the simplicity and power of hypertext

and

> htmx is small[...], dependency-free, extendable, IE11 compatible & has reduced code base sizes by 67% when compared with react

Pretty cool, huh ?! The idea of building websites without using a big JS framework and focusing back on the HTML sounds really interesting and, well, kinda new for someone who got into development after React was created and widespread.

Here is a list of what interested me :

- I am tired of big frameworks like Vue or React that need their own dev server and that can be really slow sometimes
- I also tend to dislike build processes, as they can be long or completely ignored (in a company I was previously, we used the dev server as the production server)
- We can have the frontend built alongside the backend and not in a completly different application
- Having simply HTML + a lightweight script can boost significantly the performance of web apps
- ...

## Testing HTMX

First let's go with a simple example we can find on the HTMX webpage

```html
  <script src="https://unpkg.com/htmx.org@1.9.10"></script>
  <!-- have a button POST a click via AJAX -->
  <button hx-post="/clicked" hx-swap="outerHTML">
    Click Me
  </button>
```

What this code does is, when the button is clicked, send a post request to `/clicked` and swap the button with the HTML response. One thing already caught my attention : **an HTML response ?!**

I have been used to have to write JSON APIs that I forgot there could be another way to send data to a page. JSON is great because it interfaces very nicely with Javascript, and since most of the reactive stuff in our apps is written in JS, JSON APIs look like the most sensible thing to do. However, JSON doesn't really interfaces with HTML really well. You know what does ? HTML !!! This makes a lot of sense ! And actually having the server take care of how 