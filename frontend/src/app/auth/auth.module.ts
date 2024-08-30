/**
 * @license
 * Copyright Akveo. All Rights Reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 */

import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { AuthRoutingModule } from './auth-routing.module';
import { NbAuthModule } from '@nebular/auth';
import { NbAlertModule, NbButtonModule, NbCardModule, NbCheckboxModule, NbIconComponent, NbIconModule, NbInputModule, NbLayoutComponent, NbLayoutModule, NbSidebarModule } from '@nebular/theme';

import { LoginComponent } from './login.component';
import { AuthPageComponent } from './auth.component';
import { ThemeModule } from '../@theme/theme.module';
import { SocialCallbackComponent } from './social-callback.component';


@NgModule({
    imports: [
        CommonModule,
        FormsModule,
        NbAlertModule,
        NbInputModule,
        NbButtonModule,
        NbCheckboxModule,
        AuthRoutingModule,
        ThemeModule,
        NbIconModule,
        NbLayoutModule,
        NbSidebarModule,
        NbCardModule,
    ],
    declarations: [
        AuthPageComponent,
        LoginComponent,
        SocialCallbackComponent,
    ],
})
export class AuthModule {
}