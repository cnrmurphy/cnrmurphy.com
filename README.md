# cnrmurphy.com
This repository hosts the source code for my personal website.

## Approach
Each page of my website is written in Markdown and served over both **TCP** and **HTTP**:

- The **HTTP** version is a traditional browser-based interface.
- The **TCP** server allows you to browse my site directly from the terminal.

Markdown files are modular — for example, each section of my resume lives in its own file. This allows terminal users to run commands like `contact` or `experience` to read individual sections.

To view the full resume, I simply concatenate the files using `cat`. The same modular approach is used to generate the HTML version, making the system **clean, reusable, and format-agnostic**.

## Why...
**Markdown?**

Markdown is simple, clean, and portable. It’s easy to maintain, displays well across formats, and can be converted into HTML, plain text, or even PDF with existing tools.

**Not use a static site generator?**

There are plenty of static site generators out there, but I’m not interested in adding unnecessary tooling. I don't want to deal with JavaScript, theming systems, or learning another templating language.

I wanted to use tools I already enjoy — minimal dependencies, minimal complexity, and something that feels fun and natural to me.

**TCP?**

These days, I'm feeling increasingly overwhelmed by the modern web: too much noise, too little control. I'm trying to reclaim simplicity — using tools like RSS, spending more time in the terminal, and avoiding algorithm-driven feeds.

It reminds me of when I first started using the internet as a kid back in the early 2000's.
Before even using IRC, I thought it was amazing to be able to `telnet sdf.lonestar.org`, create an account, and chat with tech-minded folks all from within the terminal.

So I thought it would be fun to provide an alternative way for people to access my website by creating a browsing experience within the terminal.

As a bonus, I can share my resume in a way that’s a little odd, but hopefully memorable:
`printf "resume\nbye\n" | nc cnrmurphy.com 2001 > resume.txt`

