import { useSettings } from '@/context/SettingsContext';
import { translations } from '@/locales';

export const useTranslation = () => {
  const { settings } = useSettings();
  const locale = settings.language.locale;
  const t = translations[locale];

  // Helper function to replace placeholders like {theme}
  const translate = (key: string, params?: Record<string, string>): string => {
    let text = key;
    if (params) {
      Object.entries(params).forEach(([param, value]) => {
        text = text.replace(`{${param}}`, value);
      });
    }
    return text;
  };

  return { t, locale, translate };
};
