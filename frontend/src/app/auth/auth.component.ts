import { Component } from '@angular/core';
import { NbLoginComponent } from '@nebular/auth';

@Component({
  selector: 'ngx-authx',
  styleUrls: ['./auth.component.scss'],
  template: `
    <nb-layout center=true windowMode=true withScroll="false" >
      <nb-layout-column >
        <router-outlet></router-outlet>
      </nb-layout-column>
    </nb-layout>


  `

})
export class AuthPageComponent { }