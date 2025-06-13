(function($){
  $.fn.dialog = function(opts){
    return this.each(function(){
      var $self = $(this);
      var placeholder = $('<span style="display:none"></span>');
      $self.after(placeholder);
      var overlay = $('<div class="dlg-overlay"></div>').css({
        position:'fixed',top:0,left:0,right:0,bottom:0,
        background:'rgba(0,0,0,0.3)',display:'flex',
        'align-items':'center','justify-content':'center'
      });
      var box = $('<div class="dlg-box"></div>').css({
        background:'#fff',padding:'1em',border:'1px solid #666'
      });
      var btns = $('<div class="dlg-buttons"></div>').css({'text-align':'right','margin-top':'1em'});
      var close = function(){ overlay.remove(); placeholder.after($self.hide()); placeholder.remove(); };
      if(opts && opts.buttons){
        Object.keys(opts.buttons).forEach(function(name){
          $('<button type="button"></button>').text(name).on('click',function(){
            close();
            opts.buttons[name].call($self[0]);
          }).appendTo(btns);
        });
      }
      box.append($self.show()).append(btns);
      overlay.append(box).appendTo('body');
      $self.data('dlg-overlay', overlay);
    });
  };
})(jQuery);
