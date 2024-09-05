import { Component, Input } from '@angular/core';
import { NbSortDirection, NbSortRequest, NbTreeGridDataSource, NbTreeGridDataSourceBuilder } from '@nebular/theme';

interface TreeNode<T> {
  data: T;
  children?: TreeNode<T>[];
  expanded?: boolean;
}

interface FSEntry {
  time: string;
  assert: string;
  sl: number;
  action: string;
  usdt:number;
}

@Component({
  selector: 'ngx-trade-history',
  templateUrl: './trade-history.component.html',
  styleUrls: ['./trade-history.component.scss'],
})
export class TradeHistoryComponent {
  customColumn = 'time';
  defaultColumns = ['assert', 'sl', 'action', 'usdt'];
  allColumns = [this.customColumn, ...this.defaultColumns];

  dataSource: NbTreeGridDataSource<FSEntry>;

  sortColumn: string;
  sortDirection: NbSortDirection = NbSortDirection.NONE;

  constructor(private dataSourceBuilder: NbTreeGridDataSourceBuilder<FSEntry>) {
    this.dataSource = this.dataSourceBuilder.create(this.data);
  }

  updateSort(sortRequest: NbSortRequest): void {
    this.sortColumn = sortRequest.column;
    this.sortDirection = sortRequest.direction;
  }

  getSortDirection(column: string): NbSortDirection {
    if (this.sortColumn === column) {
      return this.sortDirection;
    }
    return NbSortDirection.NONE;
  }

  private data: TreeNode<FSEntry>[] = [
    {
      data: { time: '01/02/24 10:00', assert: 'ETH', sl: 2, action: "buy", usdt: 100 },
      children: [
       
      ],
    },
    {
      data: { time: '01/02/24 177:00', assert: 'BTC', sl: 5, action: "sell", usdt: 55 },
      children: [
      ],
    },
    {
      data: { time: '01/02/25 07:00', assert: 'SONDQ', sl: 399, action: "buy", usdt: 100 },
      children: [
      ],
    },
  ];

  getShowOn(index: number) {
    const minWithForMultipleColumns = 400;
    const nextColumnStep = 100;
    return minWithForMultipleColumns + (nextColumnStep * index);
  }
}

