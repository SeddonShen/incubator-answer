import { FC } from 'react';
import { useTranslation } from 'react-i18next';

import { usePageTags } from '@/hooks';
import { useLegalTos } from '@/services';
import { htmlToReact } from '@/utils';

const Index: FC = () => {
  const { t } = useTranslation('translation', { keyPrefix: 'nav_menus' });
  usePageTags({
    title: t('tos'),
  });
  const { data: tos } = useLegalTos();
  const contentText = tos?.terms_of_service_original_text;
  let matchUrl: URL | undefined;
  try {
    if (contentText) {
      matchUrl = new URL(contentText);
    }
    // eslint-disable-next-line no-empty
  } catch (ex) {}
  if (matchUrl) {
    window.location.replace(matchUrl.toString());
    return null;
  }
  return (
    <div>
      <h3 className="mb-4">{t('tos')}</h3>
      <div className="fmt">
        {htmlToReact(tos?.terms_of_service_parsed_text || '')}
      </div>
    </div>
  );
};

export default Index;
