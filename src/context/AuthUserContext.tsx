import { createContext, useContext, Context } from "react";
import useFirebaseAuth from "../lib/useFirebaseAuth";
import firebase from 'firebase';

export interface AuthUserContext {
    authUser?: firebase.UserInfo | null;
    loading: boolean;
    signInWithEmailAndPassword: (email: string, password: string)=> Promise<any>;
    signOut: () => Promise<void>;

}
 
const authUserContext = createContext<AuthUserContext>({
    authUser: null,
    loading: true,
    signInWithEmailAndPassword: async(email: string, password: string): Promise<any> => {},
    signOut: async() =>{}
});

export function AuthUserProvider({children}: any) {
    const auth = useFirebaseAuth();

    return <authUserContext.Provider value={auth}>
        {children}
    </authUserContext.Provider>
}

export const useAuth = () => useContext(authUserContext);