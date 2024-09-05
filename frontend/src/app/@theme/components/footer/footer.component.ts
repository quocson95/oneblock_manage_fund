import { Component } from '@angular/core';

@Component({
  selector: 'ngx-footer',
  styleUrls: ['./footer.component.scss'],
  template: `
    <span class="created-by">
      <!-- Created with ♥ by <b><a href="https://akveo.page.link/8V2f" target="_blank">Akveo</a></b> 2019 -->
    </span>
    <div class="socials">
      <!-- <a href="#" target="_blank" class="ion ion-social-github"></a> -->
      <a href="https://www.facebook.com/groups/oneblockvn" target="_blank" class="ion ion-social-facebook"></a>
      <a href="https://x.com/oneblockvn" target="_blank" class="ion ion-social-twitter"></a>
      <a  href="https://t.me/oneblockvnchannel" target="_blank" ><nb-icon icon="message-circle"  style="font-size: 1.8rem;" pack="eva"></nb-icon></a>
       
    </div>
  `,
})
export class FooterComponent {
}
