import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})


export class AppComponent {
  showDashboard: boolean = true;
  showList: boolean = false

  setShowDashboard(show: boolean): void {
    this.showDashboard = show;
    this.setShowList(!show);
  }

  setShowList(show: boolean): void {
    this.showList = show;
  }

  onPropListVisible(show: boolean): void {
    if(show) {
      this.setShowDashboard(false);
      this.setShowList(true);
      console.log("event received - dashboard: "+this.showDashboard+" - list "+this.showList)
    }
  }
}
