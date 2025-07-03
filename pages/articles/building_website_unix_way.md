# Building My Personal Website — The Unix Way
- - - 
**June 2025**

Recently I decided it was about time to stand up a peronal website. I wanted a digital waystone of sorts to host things like personal services, articles, my resume, etc.
The first question I asked myself was:
**What’s the simplest and most enjoyable way I could build this?**
The second question I asked myself was:
**How can I make it fun?**

There are a multitude of solutions for build a website these days. I could have used a popular and familiar framework like React. But, that would certainly be overkill. And not fun!
I could use one of the many static site generators that exist. I a few are Hugo, Jekyll, and Gatsby. They also take advantage of Markdown documents as a medium for writing
pages. I enjoy working with Markdown so I would define that as fun. Furthermore, this is how I intended to work with my own personal website - I didn't want to deal too much with
actually writing out HTML or crafting my own content manager. Markdown already exists, is familiar, easy to work with, and well formatted. A suitable solution indeed. But, I just
want to get something up quick and not deal with learning one of these many static site generators. At the moment, I am not interested in the flare they offer in the form of templates
and such.

Surely I should just be able to easily convert Markdown to HTML and throw something up quick. With a little research, I found a way forward here. I could use a neat program
called `Pandoc` that allows me to do just that! It also allows you to convert from MD to quite a few other formats as well. While I don't have an immediate use for this, it is
a piece of knowledge I'm happy I now have.

## My Goals

- Minimal dependencies  
- Simple, transparent structure  
- Avoid HTML and JS as much as possible  
- Use Markdown as the source of truth  
- Reusable content across browser and terminal  
- Make it fun

## The Stack

Here's what I used:

- **Go** – for serving HTTP and TCP requests  
- **Makefile** – for basic automation  
- **Markdown** – for writing all content  
- **Pandoc** – for converting Markdown to HTML  
- **Mdcat** – for pretty-printing Markdown in the terminal

That's it.

## Serving Markdown over HTTP

The basic idea is this: all my articles, resume sections, and content are written in Markdown. I use `pandoc` to convert them to HTML. Then I serve them using a minimal Go HTTP server.

There’s no JS framework, no build pipeline, or other bloat - just raw HTML rendered from Markdown. Pages load fast, the codebase is small, and it is a system tailored to my needs.

This makes the whole site easy to manage. If I want to write something new, I create a `.md` file. If I want to publish, I run `make`.

## ...and Also Over TCP

But I didn’t stop there.

I’ve been simplifying my digital life — using RSS again, spending less time on noisy platforms, and leaning back into the terminal. It made me nostalgic for when I’d `telnet sdf.lonestar.org` from a Windows command prompt and get dropped into a shared Unix system. That kind of interaction had a charm — minimal, efficient, and personal.

So I thought:  
**What if you could browse my website entirely from the terminal?**

No web browser. Just a TCP connection and a few commands. Like:

```bash
printf "resume\nbye\n" | nc cnrmurphy.com 2003
```

This command connects to my TCP server, prints out my resume in your terminal (nicely formatted with `mdcat`), and disconnects.

You can even save it:

```bash
printf "resume\nbye\n" | nc cnrmurphy.com 2003 | tee cnr_resume.txt
```

The resume is just Markdown rendered cleanly for the terminal — no CSS, no dependencies. Just text, styled using conventions that have worked for decades.

## Reusing the Markdown

Because the Markdown files are the source of truth, I reuse them for both the web and TCP interface. For example, my resume is split into sections:

```
resume/
├── contact.md
├── experience.md
├── projects.md
└── education.md
```

Over HTTP, I use `pandoc` to render each section as a full HTML page. Over TCP, I use `mdcat` to pretty-print them in sequence, separated by horizontal rules. Bash handles the composition, and Go handles the TCP interactions.

The file structure looks like this:

```
cnrmurphy.com/
├── pages/
│   ├── resume/
│   │   ├── contact.md
│   │   ├── experience.md
│   │   ├── projects.md
│   │   └── education.md
│   ├── articles/
│   │   └── building_website_unix_way.md
│   └── about.md
├── tcp.go
├── http.go
├── main.go
├── Makefile
└── go.mod
```

Each TCP command corresponds to a Markdown file. Want to read just the `projects` section? Type `projects`. Want the full resume? Type `resume`.

You can even grab this very article the same way:

```bash
printf "articles resume_over_tcp\nbye\n" | nc cnrmurphy.com 2003 | tee resume_over_tcp.txt
```

## Final Thoughts

I didn’t want to build *just* another personal site. I wanted to build something I’d enjoy maintaining — something that reflected the way I think about software, the web, and computing in general.

This approach isn’t for everyone. But it’s fast, clean, and deeply satisfying. And to me, that’s the point.
