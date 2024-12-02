use std::fs::File;
use std::path::Path;
use std::io::{self, BufRead};

fn parse_data(filename: String) -> Vec<Vec<i32>> {
    let path = Path::new(&filename);
    let file = match File::open(path) {
        Ok(file) => file,
        Err(e) => {
            eprintln!("{}", e);
            return vec![];
        }
    };

    let reader = io::BufReader::new(file);
    let mut stuff: Vec<Vec<i32>> = Vec::new();

    for lines in reader.lines() {
        match lines {
            Ok(content) => {
                let parts: Vec<&str> = content.split_whitespace().collect();
                let mut cont: Vec<i32> = Vec::new();
                for part in parts {
                    let s: i32 = part.parse().unwrap();
                    cont.push(s);
                }

                stuff.push(cont);
            }
            Err(e) => {
                eprintln!("{}", e);
            }
        };
    }

    return stuff;
}

fn increasing(subarr: &Vec<i32>) -> bool{
    let size = subarr.len();
    let mut safe = false;

    for i in 0..size {
        if i == size-1 {
            continue
        }

        if subarr[i] == subarr[i+1]-1 || subarr[i] == subarr[i+1]-2 || subarr[i] == subarr[i+1]-3 {
            safe = true;
           }else{
            safe = false;
            break
           }
    }
    return safe;
}

fn decreasing(subarr: &Vec<i32>) -> bool{
    let size = subarr.len();
    let mut safe = false;

    for i in 0..size-1{
       if i == size-1 {
            continue
       }

       if subarr[i] == subarr[i+1]+1 || subarr[i] == subarr[i+1]+2 || subarr[i] == subarr[i+1]+3 {
        safe = true;
       }else{
        safe = false;
        break
       }
    }
    return safe;
}

fn num_of_safe(data: &Vec<Vec<i32>>) -> i32 {
    let mut total: i32 = 0;
    for subarr in data {
        let mut safe: bool = false;
        if subarr[0] <= subarr[1] {
            safe = increasing(&subarr)
        }else if subarr[0] >= subarr[1] {
            safe = decreasing(&subarr)
        }

        if safe {
            total += 1;
        }
    }


    return total;
}

fn main() {
    let data = parse_data(String::from("../data.txt"));
    let safe = num_of_safe(&data);
    println!("{}", safe);
}