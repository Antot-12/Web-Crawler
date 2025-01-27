# 🕸️ Web Crawler 🕷️

Hi everyone! 👋  
This is a **Multi-Threaded Web Crawler** written in Go! 🤓🚀

## 🧐 What does it do?  
This program **crawls the web** like a spider 🕸️:
1. Starts with a website you give it 🖥️.
2. Finds all the links on the webpage 🔗.
3. Recursively visits those links 🌐 (but only to a certain depth so it doesn't crawl forever! 😅).
4. Shows you all the crawled URLs in the terminal! 💻
 
---

## 📖 Example Output  
When you start the program, it might look like this in your terminal:  
```plaintext
Starting web crawl at https://google.com with a depth of 2
Crawling: https://google.com (Depth: 2)
Crawling: https://accounts.google.com/ServiceLogin?hl=en&passive=true&continue=https://www.google.com/&ec=GAZAAQ (Depth: 1)
Crawling: https://www.youtube.com/?tab=w1 (Depth: 1)
Crawling: https://drive.google.com/?tab=wo (Depth: 1)
Crawling: https://www.google.com/imghp?hl=en&tab=wi (Depth: 1)
Crawling: https://news.google.com/?tab=wn (Depth: 1)
Crawling: https://maps.google.com/maps?hl=en&tab=wl (Depth: 1)
Crawling: https://play.google.com/?hl=en&tab=w8 (Depth: 1)
Crawling: https://www.google.com/intl/en/about/products?tab=wh (Depth: 1)
Crawling: https://mail.google.com/mail/?tab=wm (Depth: 1)
Crawling: http://www.google.com/history/optout?hl=en (Depth: 1)
Crawling completed.
```

Pretty cool, huh? 😎

---

## 🚀 How to run?  

1. **Install Go**: First, make sure you have [Go installed](https://go.dev/) 🛠️.  
1. Clone the repository:
   ```bash
   git clone https://github.com/Antot-12/Web-Crawler.git

2. Navigate to the project folder:
   ```bash
   cd Web-Crawler
   ```
   
4. **Run it**: Use this command:  
   ```bash
   go run crawler.go
   ```  
4. **Customize it**: Change the `startURL` and `maxDepth` values in the code to crawl other websites or limit how deep it goes 🧗.  

---
