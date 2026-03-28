extern crate slug;
use slug::slugify;
use std::io::{self, Write};

fn prompt(label: &str) -> String {
    print!("{}: ", label);
    io::stdout().flush().unwrap();
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();
    input.trim().to_string()
}

fn find_site_root() -> std::path::PathBuf {
    let mut dir = std::env::current_dir().unwrap();
    loop {
        if dir.join("config.toml").exists() {
            return dir;
        }
        if !dir.pop() {
            panic!("Could not find Zola site root (no config.toml found)");
        }
    }
}

pub fn run() {
    let title = prompt("Talk title");
    let speaker = prompt("Speaker name (e.g. Jane Doe)");

    let slug = slugify(title.to_lowercase());
    let root = find_site_root();
    let filename = root.join(format!("content/talks/{slug}.md"));

    let speaker_slug = slugify(&speaker);
    let content = format!(
        r#"+++
title = "{title}"

[extra]
speaker = "{speaker_slug}"
+++
"#
    );

    // Create speaker if doesn't exist
    let speaker_file = root.join(format!("content/speakers/{speaker_slug}.md"));
    if !speaker_file.exists() {
        let speaker_content = format!(
            r#"+++
title = "{speaker}"
+++
"#
        );
        std::fs::write(&speaker_file, speaker_content).unwrap();
        println!("✏️ Created new speaker {}", speaker_file.display());
    } else {
        println!("This is a returning speaker 🎉")
    }

    std::fs::write(&filename, content).unwrap();
    println!("✏️ Created new talk {}", filename.display());
}
