import { savePostFetcher } from "@/features/apis/savePost";
import { validateFormEventTarget } from "@/utils/form";
import { FormEvent, useCallback, useState } from "react";
import { toast } from 'sonner';
import { useSWRConfig } from "swr"

export const usePostForm = () => {
    const USER_ID = 1;
    const [isFocus, setIsFocus] = useState(false);
    const { mutate } = useSWRConfig();

    const handleOnSubmit = (onClose?: () => void) => useCallback(async (e: FormEvent<HTMLFormElement>) => {
      const { formData, form } = validateFormEventTarget(e.target);
      const { title, body } = Object.fromEntries(
        [...formData.entries()].filter((v): v is [string, string] => typeof v[1] === 'string'),
      );

      e.preventDefault();
      try {
        await savePostFetcher(USER_ID, title, body);
        await mutate("/tweets");
        form.reset();
        toast('投稿を作成しました！');
        if (onClose == undefined) {
          return
        }
        onClose();
      } catch (error) {
        console.error('Error:', error);
      }
    }, []);

    return {
        handleOnSubmit,
        isFocus,
        setIsFocus
      };
}