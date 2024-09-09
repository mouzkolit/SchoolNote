import { writable } from 'svelte/store';

interface SchoolStore {
    schoolName: string;
    schoolId: string;
}

export const schoolStore = writable<SchoolStore>({
    schoolName: '',
    schoolId: ''
});