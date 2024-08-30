import { AfterViewInit, Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { NbAuthService } from '@nebular/auth';


@Component({
    selector: 'ngx-social-callback',
    templateUrl: './social-callback.component.html'
})
export class SocialCallbackComponent implements OnInit {
    show = false;
    loginSuccess = false;
    constructor(
        private activeRoute: ActivatedRoute,
        private auth: NbAuthService,
        private router: Router,
    ) { }


    ngOnInit(): void {
        this.activeRoute.queryParams.subscribe(params => {
            console.log(params);
            const id = params["id"];
            if (!!id && id.length > 0) {
                this.loginSuccess = true;
                this.auth.authenticate("google-sso", id).subscribe(result => console.log(result));
                setInterval(() => {
                    this.router.navigate(['/']);
                }, 1000);
            }
            this.show = true;
        })

    }
}