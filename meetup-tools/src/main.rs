mod commands;

fn main() {
    println!("Hello, world!");
    let args: Vec<String> = std::env::args().collect();
    match args.get(1).map(|s| s.as_str()) {
        Some("new-talk") => commands::new_talk::run(),
        _ => eprintln!("Usage: meetup [new-event|new-talk|relations]"),
    }
}
