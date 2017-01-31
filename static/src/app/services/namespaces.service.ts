/*** Service for Get Namespaces ***/

//import
import { Injectable }     from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import {Observable} from 'rxjs/Rx';

// Import RxJs required methods
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';

@Injectable()
export class NamespacesService {
     // Resolve HTTP using the constructor
     constructor (private http: Http) {}
     // private instance variable to hold base url
     private NamespacesUrl = '/api/namespaces'; 

     getNamespaces() : Observable<string[]> {

         console.log("faccio la chiamata")

         return this.http.get(this.NamespacesUrl)
                .map((res:Response) => res.json().data)
                .catch((error:any) => Observable.throw(error.json().error || 'Server error'))

     }
}

