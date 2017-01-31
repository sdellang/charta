/*** Service to get Properties ***/

//import
import { Injectable }     from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import { Observable } from 'rxjs/Rx';
import { Pod } from '../model/pod'

// Import RxJs required methods
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';

@Injectable()
export class PropertiesService {
     // Resolve HTTP using the constructor
     constructor (private http: Http) {}
     // private instance variable to hold base url
     private NamespacesUrl = '/api/namespaces/'; 

     getProperties(namespace: string) : Observable<Pod[]> {

         console.log("getting properties for:"+namespace)

         return this.http.get(this.NamespacesUrl+namespace)
                .map((res:Response) => res.json().data)
                .catch((error:any) => Observable.throw(error.json().error || 'Server error'))

     }
}