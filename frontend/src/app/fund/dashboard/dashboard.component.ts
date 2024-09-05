import { Component, OnDestroy } from '@angular/core';
import { NbThemeService } from '@nebular/theme';
import { takeWhile } from 'rxjs/operators';
import { SolarData } from '../../@core/data/solar';

interface CardSettings {
    title: string;
    iconClass: string;
    type: string;
}

@Component({
    selector: 'ngx-fund-dashboard',
    styleUrls: ['./dashboard.component.scss'],
    templateUrl: './dashboard.component.html',
})
export class FundDashboardComponent implements OnDestroy {
    constructor(private themeService: NbThemeService,
        private solarService: SolarData) { }
    ngOnDestroy(): void {

    }
}