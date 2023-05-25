use std::time::Duration;

use chromiumoxide::{cdp::js_protocol::runtime::EvaluateParams, Browser, BrowserConfig};
use futures::StreamExt;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let (mut browser, mut handler) =
        Browser::launch(BrowserConfig::builder().with_head().build()?).await?;

    let handle = tokio::spawn(async move {
        while let Some(h) = handler.next().await {
            if h.is_err() {
                println!("handle err: {:?}", h);
                break;
            }
        }
    });

    let pages = browser.pages().await?;
    println!("{:?}", pages);

    let page = browser.new_page("https://baidu.com").await?;
    let page = page.wait_for_navigation().await?;

    tokio::time::sleep(Duration::from_secs(10)).await;
    let eval = EvaluateParams::builder().expression("() => {return 123;}");
    let result = page.evaluate(eval.clone().build()?).await?;
    println!("{:?}", result.into_value::<usize>());

    browser.close().await?;
    handle.await?;
    Ok(())
}
