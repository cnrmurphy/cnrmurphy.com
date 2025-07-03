# cnrmurphy.com
This repository hosts the source code for my personal website.

## Approach
Each page of my website is written in Markdown and served over both **TCP** and **HTTP**:

- The **HTTP** version is a traditional browser-based interface.
- The **TCP** server allows you to browse my site directly from the terminal.

I used `Pandoc` to generate a simple HTML wrapper and then use it to convert all of the Markdown pages into HTML files. A simple shell script uses `awk` to
to do template substitution. Areas to be subsituted are marked with HTML comments. This also means we can resuse our pages for the TCP server. But, in the case
of TCP, we use `mdcat` to create a pretty printed version of the page that is terminal friendly.

Finally, we use `Make` to tie everything together - compiling our files, distributing static files, running our dev server, and deploying to production.

## Development

To run the site locally for development:

```bash
# Generate HTML files from Markdown
make generate

# Start both HTTP (port 8080) and TCP (port 2003) servers
make dev
```

Visit `http://localhost:8080` for the web interface or connect to the TCP server:
```bash
nc localhost 2003
```

## Production Deployment

For production deployment using nginx and systemd:

```bash
# Pull latest changes
git pull

# Deploy everything: build binary, generate HTML, configure nginx, restart services
make deploy
```

This will:
- Generate HTML files from Markdown sources
- Build and install the Go binary for the TCP server
- Install nginx configuration for HTTP serving
- Restart both the TCP server (systemd) and nginx

## Why...
**Markdown?**

I like markdown. It's simple and I can edit it in vim. It converts to HTML easily and I don't need to think about styling while writing.

**Not use a static site generator?**

I am not familiar with these and don't feel like taking the time to learn them. This approach for me works as it uses tools I am familiar with and is intuitive. I believe
there is a lot of crossover with how those frameworks work, but this way I avoid dependencies that I don't want and have a more controlled approach.

**TCP?**

This is a bit of an ode to my first experiences of using the internet as a young lad. I'm fond of the times when I was beginning to learn program and chat with folks on `sdf.lonestar.org`.


These days, I'm feeling increasingly overwhelmed by the modern web. Ads everywhere, noisy social media networks, bloated websites. I find myself moving away from these things
and becoming ever reliant on my terminal; A sort of bunker if you will. So I thought it woud be neat if I could essentially host my website through TCP. I also like the idea
of someone being able to grab my resume from their terminal with some simple unix commands:

`printf "resume\nbye\n" | nc cnrmurphy.com 2001 | tee conor_murphy_resume.txt`

