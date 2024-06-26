import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import i18n from "i18next";
import detector from "i18next-browser-languagedetector";
import backend from "i18next-http-backend";
import {initReactI18next} from "react-i18next";
import { loader } from '@monaco-editor/react';
import * as monaco from 'monaco-editor';
import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker';
// https://github.com/remcohaszing/monaco-yaml/issues/150
import yamlWorker from '@/utils/workaround-yaml.worker?worker'

i18n
    .use(detector)
    .use(backend)
    .use(initReactI18next) // passes i18n down to react-i18next
    .init({
        debug: true,
        backend: {
            loadPath: '/i18n/{{lng}}/{{ns}}.json',
            addPath: '/i18n/{{lng}}/{{ns}}.json',
        },
        // lng: "en", // if you're using a language detector, do not define the lng option
        fallbackLng: "zh-CN",

        interpolation: {
            escapeValue: false // react already safes from xss => https://www.i18next.com/translation-function/interpolation#unescape
        },
        saveMissing: true, // send not translated keys to endpoint,
    });

window.MonacoEnvironment = {
    getWorker(_, label) {
        if (label === 'yaml') {
            return new yamlWorker();
        }
        return new editorWorker();
    },
};
loader.config({ monaco });

ReactDOM
    .createRoot(document.getElementById('root')!)
    .render(
        <React.StrictMode>
            <App/>
        </React.StrictMode>,
    )