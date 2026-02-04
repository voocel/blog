import { en } from '@/locales/en';
import { zh } from '@/locales/zh';

export type Locale = 'en' | 'zh';

export const translations = {
  en,
  zh,
};

export const localeNames: Record<Locale, string> = {
  en: 'English',
  zh: '中文',
};

export const defaultLocale: Locale = 'en';
