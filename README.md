# glna
This should eventually become a small DLNA server written in go that supports native Filesystem Monitoring on Windows, Linux and BSD.
I also use this to learn go. So it is possible (and likely) that you will find badly written or inefficient code in here. I will gradually improve this until I reached a level of skill and code quality that I think is publishable. 

## Roadmap
* Implement filesystem monitor (based on fs-notify)
* Implement media Scanner (based on taglib)
* Implement database backend
* Implement DLNA service 

The initial focus will be on Audio files. Video files will follow later.
