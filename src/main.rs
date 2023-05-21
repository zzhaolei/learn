use tokio::sync::oneshot;

#[tokio::main]
async fn main() {
    let (tx1, rx1) = oneshot::channel::<i32>();
    let (tx2, rx2) = oneshot::channel::<i32>();

    tokio::spawn(async {
        let _ = tx1.send(1);
    });

    tokio::spawn(async {
        let _ = tx2.send(2);
    });

    tokio::select! {
        val = rx1 => {
            println!("Recv: {:?}", val);
        },
        val = rx2 => {
            println!("Recv: {:?}", val);
        }
    }
}
