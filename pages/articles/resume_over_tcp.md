# Resume Over TCP
6/05/2025

I figured it was well about time to stand up my own personal website; a place to share my thoughts through articles, display my personal projects, and host my resume.
A common task that I have done in the past, but ultimately neglected. So the first question that comes to mind is "what tools should I use for this task?". I want to strike
a balance between simplicity and fun. It may be tempting to just use a framework like React, but that is overkill. Then there are a myriad of static site generators to use
which are more appropraite for this task - this checks the box of "simple", but not quite "fun". I think I can achieve something even more simple that I would find fun.
So let's define "fun". First, I think it would be fun to create a website using as little HTML and JavaScript as possible. Second, I would like to use Markdown files.
Markdown is simple, concise, and well formatted. It renders well through a multitude of mediums and can be easily converted to various formats - including HTML.
Finally, I want to use the tools I enjoy the most. Markdown is a tool I want to use to structure my articles. I want to use vim as me editor. I want to use native Unix tools,
those that have been tried and tested, to help me navigate this whole process.

The primary motivation here is to use tools I like, in a minimal way, avoiding the bloat of modern software development. In fact, this is how I have recently approaching
my use of computers and the internet recently; keeping my setup minimal, personalized, and free of bloat. I barely find myself on social media apps due to the proliferation
of noise. Even things like Youtube I try to avoid because of the amount of "content" that I don't care about that is forced into my feed. I'm even beholden to how someone
has styled their website. This has motivated me to use things like RSS...I can get the information I need and want right from my terminal. It's nothing new and has made
me nostalgic for my younger years when I first discovered the internet. Back when I woudld `telnet sdf.lonestar.org` from my Windows commandline and have my own space
on a shared Unix machine.

So I thought, "wouldn't it be cool if I could just serve my website through the terminal?". Specifically, I think it would be neat if you could just avoid a web browser altogether,
and connect to my program, provide a simple `resume` command, and get a pretty printed markdown document right in your terminal. This is perfect because it is lightweight, you also
get to use great Unix tools that are likely already on your machine, and everything is sufficiently stylized without having to use any CSS! So now I have an idea I appreciate and think
is fun to puresure, I have the toolkit of my choice, and have achieved the minimal design and approach I desired.

The next step is to spin up a simple TCP server in Go that accepts connections over a port of our choice. We pass that connection through a handler which will read
user input and match it against a set of commands. For instance, a user can input `resume` into the program and the program will simply load `resume.md` into memory
and print it to Stdout. But we want to pretty print the text. We could write our own tool to do this, but there exists a perfectly good tool out there already call mdcat!
Mdcat is short for Markdown Cat. You are probably familiar with the Unix program `cat`. If not, it is a program that concatenates multiple files together and writes the result
to stdout. A common use of cat is not actually its intended purpose, but to simply output the contents of a file to stdout. Similarly, mdcat will concatenate multiple markdown
documents and pretty print them to stdout such that they are easy to read. So when a user requests my resume, the program will just make a call to the system using `os.Execute` and
run `md resume.md` returning my nicely formatted resume.

I thought about this for a bit and realized there is some overlap between my resume and what I would like to be shown on my personal website. My resume is structured as follows:

```
resume/
├── contact
├── experience
├── projects
└── education
```

So it is pretty natural to then reimagine each section of my resume as an individual page *and* and individual command. A bit redundent perhaps, but hey, maybe
someone is just interested in seeing my projects. Therefore the following structure arises (I'll share the whole server structure to make things clear)

```
cnrmurphy.com/
├── pages/
│   ├── resume
│   │   ├── contact.md
│   │   ├── experience.md
│   │   ├── projects.md
│   │   └── education.md
│   └── about.md
├── README.md
├── tcp.go
├── http.go
├── main.go
└── go.mod
```

Great, now each individual command can be mapped to a respective markdown file and we can use `mdcat` to render it! And we have reusability this way too - when 
a user provides the `resume` command, our program will call out to `mdcat contact.md experience.md projects.md education.md` to get a perfectly reconstructed resume
in te terminal. In practice, I use some extra bash scripting to separate each section with a horizontal ruler by insert the `- - - ` syntax after each concatination.

and now you can grab my resume in your terminal and save it directly to your machine simply by running `printf "resume\nbye\n" | nc cnrmurphy.com 2003 | tee cnr_resume.txt`!
All of these shell commands are programs that already exist on Unix like machines. It connects the program running on my server, outputs the commands `resume` and `bye` to first
request the resume to be printed to stdout and then disconnect from the server, and finally it saves the output to `cnr_resume.txt` and prints the contents in your own terminal.

You can do the very same thing for this article!
`printf "articles resume_over_tcp\nbye\n" | nc cnrmurphy.com 2003 | tee resume_over_tcp.txt`
