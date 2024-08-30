/**
 * @license
 * Copyright Akveo. All Rights Reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 */
import { Component } from '@angular/core';
import { NbLoginComponent } from '@nebular/auth';

@Component({
  selector: 'ngx-login',
  // styleUrls:['./login.component.css'],
  templateUrl: './login.component.html',
})
export class LoginComponent extends NbLoginComponent {

  login() {
    console.log('login');
    window.open('https://fund.oneblock.vn/api/auth/google_signin', '_self')
  }

}
