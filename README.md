Zwitscher - mehr als nur piep piep!
===================================
![zwitscher](https://github.com/mogoh/Zwitscher/raw/master/zwitscher.png "[t͡svɪtʃɹ]")

Zwitscher wird ein neuer Twitterclient, geschrieben in Go und GTK.

Zwitscher is going to be a new Twitter client, written in Go and GTK.

Compiling and Installing
------------------------


  * [Install](http://golang.org/doc/install.html) Go as described

  * Update Go to tip (notice the update tip, instead of release):

    $ cd go/src  
    $ hg pull  
    $ hg update tip  
    $ ./all.bash  
    
  * Install iconv: goinstall github.com/sloonz/go-iconv/src
  
  * Install twister/oauth: goinstall github.com/garyburd/twister/oauth

  * Install go-gtk (don't forget to install gtk development packages before):

    $ goinstall github.com/mattn/go-gtk/gtk

  * Download and Install Zwitscher:

    $ git clone https://github.com/mogoh/Zwitscher  
    $ cd Zwitscher  
    $ gomake install



Version
-------

Version: 0 ("Zero")
