import { Button, Divider, Flex, HStack, Icon, VStack } from '@yamada-ui/react';
import { IconType } from 'react-icons';
import { FaHouse } from 'react-icons/fa6';
import { FaUser } from 'react-icons/fa6';

type SidebarProps = {
  children: React.ReactNode;
};

type SidebarItem = {
  title: string;
  icon: IconType;
  action: () => void;
};

const items: SidebarItem[] = [
  {
    title: 'ホーム',
    icon: FaHouse,
    action: () => console.log('ホーム'),
  },
  {
    title: 'プロフィール',
    icon: FaUser,
    action: () => console.log('プロフィール'),
  },
];

export default function Sidebar({ children }: SidebarProps) {
  return (
    <Flex flexDirection="row">
      <VStack w="300px" h="100vh" padding="md" position="sticky" top="0">
        {items.map((item, index) => (
          <Button
            key={index}
            variant="ghost"
            fontSize="lg"
            leftIcon={<Icon as={item.icon} />}
            p="md"
            rounded="40"
            w="fit-content"
            onClick={item.action}
          >
            {item.title}
          </Button>
        ))}
        <Button fontSize="lg" p="md" rounded="40" bg="sky.400">
          ポストする
        </Button>
      </VStack>
      <Divider orientation="vertical" h="100vh" position="sticky" top="0" />
      {children}
    </Flex>
  );
}