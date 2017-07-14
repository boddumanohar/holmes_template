You can just create a wrapper around your library. Or could just do some thing out of the box with the existing code.

Making direct process calls will be costly so we dont recomend that way. If your library is written in a language which is good for low level programming language (ex: C, C++, Rust etc..), then we recomend using API libraries and create a go-wrapper around it. If your library your library is written in high level programming language (like python, ruby etc..)

The goal of this Template is to simply creating Holmes Totem Services. At the core, this Template should be acessable enough to anyone who wants to create a service for Holmes Totem. 