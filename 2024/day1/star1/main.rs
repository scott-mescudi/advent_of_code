
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

struct Pairs {
    left: Vec<i32>,
    right: Vec<i32>,
}

fn parse_data(filename: String) -> Pairs {
    let path = Path::new(&filename);
    let file = match File::open(path) {
        Ok(file) => file,
        Err(e) => {
            eprintln!("Error opening file: {}", e);
            return Pairs { left: vec![], right: vec![] };
        }
    };

    let reader = io::BufReader::new(file);

    let mut lnums: Vec<i32> = Vec::new();
    let mut rnums: Vec<i32> = Vec::new();

    for line in reader.lines(){
        match line {
            Ok(content) => {
                let parts: Vec<&str> = content.split_whitespace().collect();
                let value1: i32 = parts[0].parse().unwrap();
                let value2: i32 = parts[1].parse().unwrap();

                lnums.push(value1);
                rnums.push(value2);
            }
            Err(e) => {
                eprintln!("{}", e);
            }
        }
    }

    lnums.sort();
    rnums.sort();

    return Pairs { left: lnums, right: rnums };
}

fn main() {
    let data = parse_data(String::from("../data.txt"));

    let len = data.left.len();

    let mut total: i32 = 0;
    for i in 0..len {
        let res;
        if data.left[i] > data.right[i] {
            res = data.left[i] - data.right[i];
        }else {
            res = data.right[i] - data.left[i];
        }
        
        total += res;
    }
    println!("{}", total)
}
