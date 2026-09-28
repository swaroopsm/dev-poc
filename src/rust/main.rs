use std::io::{BufRead, BufReader, Write};
use std::net::{TcpListener, TcpStream};

fn main() -> std::io::Result<()> {
    let listener = TcpListener::bind("0.0.0.0:8080")?;
    println!("listening on 0.0.0.0:8080");

    for stream in listener.incoming() {
        handle(stream?)?;
    }
    Ok(())
}

fn handle(mut stream: TcpStream) -> std::io::Result<()> {
    let request_line = BufReader::new(&stream)
        .lines()
        .next()
        .transpose()?
        .unwrap_or_default();

    let (status, body) = match request_line.split_whitespace().nth(1) {
        Some("/ping") => ("200 OK", "pong"),
        Some("/") => ("200 OK", "hello world from Rust"),
        _ => ("404 NOT FOUND", "not found"),
    };

    let response = format!(
        "HTTP/1.1 {status}\r\nContent-Length: {}\r\nContent-Type: text/plain\r\n\r\n{body}",
        body.len()
    );
    stream.write_all(response.as_bytes())
}
