import { savePostFetcher } from "@/features/apis/savePost";
import { validateFormEventTarget } from "@/utils/form";
import { ChangeEvent, FormEvent, useCallback } from "react";

export const usePostForm = () => {
    const formData = new FormData();
    const USER_ID = 1;

    const handleTitleTextAreaChange = useCallback((e: ChangeEvent<HTMLTextAreaElement>) => {
        formData.append("title", e.target.value);
    }, []);

    const handleBodyTextAreaChange = useCallback((e: ChangeEvent<HTMLTextAreaElement>) => {
        formData.append("body", e.target.value);
    }, []);

    const handleOnSubmit = useCallback(async (e: FormEvent<HTMLFormElement>) => {
        const { formData, form } = validateFormEventTarget(e.target)
        const { title, body } = Object.fromEntries(
            [...formData.entries()].filter((v): v is [string, string] => typeof v[1] === 'string'),
        );

        e.preventDefault();
        try {
            savePostFetcher(USER_ID, title, body)
            form.reset();
        } catch (error) {
            console.error('Error:', error);
        }
    }, []);

    return {
        handleTitleTextAreaChange,
        handleBodyTextAreaChange,
        handleOnSubmit,
      };
}