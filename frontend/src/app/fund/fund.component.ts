import { Component } from '@angular/core';
import { NbMenuItem } from '@nebular/theme';
import { MENU_ITEMS } from './fund-menu';


@Component({
    selector: 'ngx-fund',
    styleUrls: ['fund.component.scss'],
    template: `
    <ngx-one-column-layout>
        <nb-menu [items]="menu"></nb-menu>
      <router-outlet></router-outlet>
    </ngx-one-column-layout>
  `,
})

export class FundComponent {
  menu = MENU_ITEMS;
}
