import { createContext } from 'react';
import BaseTerminal from './base.ts';
import '@xterm/xterm/css/xterm.css';

export const TerminalContext = createContext<{
  terminal: BaseTerminal | null;
}>({
  terminal: null,
});
export { default as ContainerTerminal } from './container';
export { default as TtydTerminal } from './ttyd';
export type { FlowControl, Preferences, Command } from './ttyd';
export { default as BaseTerminal } from './base';
export type { ITerminalOptions, ITheme } from '@xterm/xterm';
export type {
  AddonType,
  AddonInfo,
  ClientOptions,
  BaseTerminalOptions,
  RendererType,
} from './typing.d.ts';
